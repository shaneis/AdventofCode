#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path,

    [switch] $ShowBoards
)

. $PSScriptRoot/Show-BingoBoard.ps1
. $PSScriptRoot/Test-Row.ps1
. $PSScriptRoot/Test-Column.ps1

$PuzzleInput = Get-Content -Path $Path -Raw

[int[]] $DrawnNumbers = ($PuzzleInput -split '\r?\n')[0] -split ','

$BoardObjects = $PuzzleInput -split '\n\n' |
Select-Object -Skip 1 |
ForEach-Object -Begin { $BoardID = 0 } -Process {
    [PSCustomObject] @{
        BoardID    = $BoardID++
        Board      = $_
        IsComplete = $false
    }
}

$CompleteBoards = [Collections.Generic.List[psobject]]::new()
do {
    # Need at least 5 numbers to complete a board
    foreach ($End in 5..$DrawnNumbers.Count) {
        foreach ($BoardObject in $BoardObjects | Where-Object { -not $_.IsComplete }) {
            
            $Params = @{Board = $BoardObject.Board; CalledNumbers = $DrawnNumbers[0..$End] }
            
            $RowColumnComplete = $false
            foreach ($RC in 1..5) {
                if ($RowColumnComplete) { continue }

                $Row = Test-Row @Params -Row $RC -Verbose:$false

                if (-not $Row.IsComplete) {
                    $Column = Test-Column @Params -Column $RC -Verbose:$false
                }

                if ($Row.IsComplete -or $Column.IsComplete) {
                    $FinalNumbers = $DrawnNumbers[0..$End]
                    $FinalBoardID = $BoardObject.BoardID
                    $BoardObject.IsComplete = $true


                    if ($ShowBoards) {
                        $ShowBoard = @{
                            Board         = $BoardObject.Board
                            CalledNumbers = $FinalNumbers
                            Verbose       = $false
                        }
                        "BoardID: $FinalBoardID"
                        Show-BingoBoard @ShowBoard
                        Write-Host "$([Environment]::NewLine)"
                    }

                    Write-Verbose "Adding Board: $($BoardObject.BoardID) to complete list"
                    $CompleteBoards.Add($BoardObject)

                    $RowColumnComplete = $true
                }
            }
        }
    }
} while ($CompleteBoards.Count -ne $BoardObjects.Count)

$WinningBoard = ($BoardObjects | Where-Object BoardID -eq $FinalBoardID).Board

[int[]] $BoardNumbers = $WinningBoard -split '\r?\n' -split ' ' |
Where-Object { $_ -as [int] }

[int[]] $UnmatchedNumbers = foreach ($Number in $BoardNumbers) {
    if ($Number -notin $FinalNumbers) {
        $Number
    }
}

[int] $SumUnmatched = ($UnmatchedNumbers | Measure-Object -Sum).Sum

[int] $LastNumberCalled = $FinalNumbers[-1]

$FinalScore = $SumUnmatched * $LastNumberCalled

@'
  Board Numbers:            {0}
  Unmatched Numbers:        {1}
  Sum of unmatched numbers: {2}
  Last number called:       {3}
  =============================
  Final Score:              {4}
'@ -f @(
    ($BoardNumbers -join ',')
    ($UnmatchedNumbers -join ',')
    $SumUnmatched
    $LastNumberCalled
    $FinalScore
)
