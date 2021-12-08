#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

$PuzzleInput = Get-Content -Path $Path -Raw

[int[]] $DrawnNumbers = ($PuzzleInput -split '\n\n')[0] -split ','

Write-Verbose "Drawn numbers:$([Environment]::Newline)$($DrawnNumbers | Out-String)"

$Boards = $PuzzleInput -split '\n\n' |
    Select-Object -Skip 1

foreach ($Board in $Boards) {
    "Board:"
    $Board
}
