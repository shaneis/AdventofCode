#! /usr/bin/pwsh

[CmdletBinding()]
param(
    $Path
)

$LanternFish = Get-Content -Path $Path

$Fish = [Collections.Generic.List[int]]::new()
foreach ($Num in $LanternFish -split ',') { $Fish.Add($Num) }

foreach ($Number in 1..80) {
    $InitialCount = $Fish.Count

    for ($i = 0; $i -lt $InitialCount; $i++) {
        $NewNum = $Fish[$i] - 1

        if ($NewNum -lt 0) {
            $Fish[$i] = 6
            $Fish.Add(8)
        }
        else {
            $Fish[$i] = $NewNum
        }
    }
    [PSCustomObject] @{
        Day   = $Number
        Count = $Fish.Count
    }
}

@'
Final Count: {0}
'@ -f $Fish.Count
