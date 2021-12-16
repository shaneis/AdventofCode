#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

. $PSScriptRoot/Show-BingoBoard.ps1
. $PSScriptRoot/Test-Row.ps1
. $PSScriptRoot/Test-Column.ps1

$PuzzleInput = Get-Content -Path $Path -Raw

[int[]] $DrawnNumbers = ($PuzzleInput -split '\n\n')[0] -split ','

$Boards = $PuzzleInput -split '\n\n' |
    Select-Object -Skip 1

$BoardID = 0
$BoardObjects = foreach ($Board in $Boards) {
	[PSCustomObject] @{
		BoardID = $BoardID++
		Board = $Board
		IsComplete = $false
	}
}

$Complete = $false
do {
		# Need at least 5 numbers to have a match
		foreach ($end in 5..$DrawnNumbers.Count) {
			foreach ($Board in $BoardObjects | Where-Object {-not $_.IsComplete}) {
				Write-Verbose "Checking Board: $($Board.BoardID) for $end out of $($DrawnNumbers.Count)"

				$Params = @{
						Board = $Board.Board
						CalledNumbers = $DrawnNumbers[0..$end]
					}
				foreach ($rc in 1..5) {
					
					$Row = Test-Row @Params -Row $rc -Verbose:$false
					$Column = Test-Column @Params -Column $rc -Verbose:$false

					if ($Row.IsComplete -or $Column.IsComplete -and -not $Complete) {
						$FinalNumbers = $DrawnNumbers[0..$end]
						$FinalNumbers -join ','
						"Board: $($Board.BoardID)"
						Show-BingoBoard -Board $Board.Board -CalledNumber $DrawnNumbers[0..$end] -Verbose:$false
						$Complete = $true
						$Board.IsComplete = $true
					}
					
					if ($Complete) {break}
				}
				if ($Complete) {break}
			}
			if ($Complete) {break}
		}
} while (-not $Complete)

$WinningBoard = $Boards[$BoardID-1]

[int[]] $BoardNumbers = $WinningBoard -split '\r?\n' -split ' ' | Where-Object {$_ -as [int]}

[int[]] $UnmatchedNumbers = foreach ($Nbr in $BoardNumbers) {
    if ($Nbr -notin $FinalNumbers) {
		$Nbr
	}
}

"Board Numbers: $($BoardNumbers -join ',')"
"Unmatched: $($UnmatchedNumbers -join ',')"

$SumUnmatched = ($UnmatchedNumbers | Measure-Object -Sum).Sum
"Sum of unmatched: $SumUnmatched"

$LastNbrCalled = $FinalNumbers[-1]
"Last number called: $LastNbrCalled"

$FinalScore = $SumUnmatched * $LastNbrCalled
"Final Score: $FinalScore"

