[CmdletBinding()]
param (
    [Parameter()]
    [string]
    $Path
)

$InputContent = Get-Content -Path $Path

$StartPosition = [PSCustomObject]@{
    CounterID = 0
    Direction = 90
    X = 0
    Y = 0
    Action = ''
    Amount = 0
}

foreach ($Instruction in $InputContent) {
    $Action = $Instruction.Substring(0, 1)
    [int]$Amount = $Instruction.Substring(1)

    if ($Action -eq 'N') {
        $StartPosition.Y += $Amount
    }

    if ($Action -eq 'S') {
        $StartPosition.Y -= $Amount
    }

    if ($Action -eq 'E') {
        $StartPosition.X += $Amount
    }

    if ($Action -eq 'W') {
        $StartPosition.X -= $Amount
    }

    if ($Action -eq 'F') {
        switch ($StartPosition.Direction) {
            0 { $StartPosition.Y += $Amount }
            90 { $StartPosition.X += $Amount }
            180 { $StartPosition.Y -= $Amount }
            270 { $StartPosition.X -= $Amount }
        }
    }

    if ($Action -eq 'R') {
        $StartPosition.Direction += $Amount
    }

    if ($Action -eq 'L') {
        $StartPosition.Direction += (360 - $Amount)
    }

    $StartPosition.CounterID++
    $StartPosition.Direction %= 360
    $StartPosition.Action = $Action
    $StartPosition.Amount = $Amount

    $StartPosition
}

$StartPosition | Select-Object -Last 1 -Property *,
    @{
        Name = 'ManhattanDistance'
        Expression = { [Math]::Abs($_.X) + [Math]::Abs($_.Y) }
    }