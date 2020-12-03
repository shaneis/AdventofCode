[bool]$Test = $false

$InputText = if ($test) {
    Get-Content -Path "$($PSScriptRoot)\day03_part01_Sample.txt"
} else {
    Get-Content -Path "$($PSScriptRoot)\day03_part01.txt"
}

$startingPosition = [PSCustomObject]@{
    X = 1
    Y = 1
}

$hitATree = 0
do {
    $actualX = ($startingPosition.X % $InputText[0].Length) - 1
    $thisLine = $InputText[$startingPosition.Y - 1].ToCharArray()
    $hitOrMiss = if ($thisLine[$actualX] -eq '#') { 'X' } else { '0' }
    $thisLine[$actualX] = $hitOrMiss

    if ($hitOrMiss -eq 'X') { $hitATree++ }

    [PSCustomObject]@{
        LineNumber = $startingPosition.Y
        RowNumber = $startingPosition.X
        Line = -join $thisLine
        NumberOfConcussions = $hitATree
    }

    $startingPosition.Y += 1
    $startingPosition.X += 3
} while ($startingPosition.Y -le ($InputText).Count)
