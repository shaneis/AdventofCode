[bool]$Test = $false

$InputText = if ($test) {
    Get-Content -Path "$($PSScriptRoot)\day03_part01_Sample.txt"
} else {
    Get-Content -Path "$($PSScriptRoot)\day03_part01.txt"
}

<#

    Right 1, down 1.
    Right 3, down 1. (This is the slope you already checked.)
    Right 5, down 1.
    Right 7, down 1.
    Right 1, down 2.

#>

$all = foreach ($case in @(
    [PSCustomObject]@{ Right = 1; Down = 1 }
    [PSCustomObject]@{ Right = 3; Down = 1 }
    [PSCustomObject]@{ Right = 5; Down = 1 }
    [PSCustomObject]@{ Right = 7; Down = 1 }
    [PSCustomObject]@{ Right = 1; Down = 2}

)) {
    $startingPosition = [PSCustomObject]@{
        X = 1
        Y = 1
        Right = $case.Right
        Down = $case.Down
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
            Right = $startingPosition.Right
            Down = $startingPosition.Down
        }

        $startingPosition.Y += $startingPosition.Down
        $startingPosition.X += $startingPosition.Right
    } while ($startingPosition.Y -le ($InputText).Count)
}

$all | Where-Object LineNumber -eq $maxLine.Maximum | ForEach-Object -Begin { $total = 1 } -Process {
    $total *= $PSItem.NumberOfConcussions
} -End { $total }

