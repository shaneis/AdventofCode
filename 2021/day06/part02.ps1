#! /usr/bin/pwsh

[CmdletBinding()]
param(
    $Path,

    $Days
)

$LanternFish = Get-Content -Path $Path

$LanternFish -split ',' | ForEach-Object -Begin {
    $Fish = @{}
} -Process {
    if (-not $Fish.ContainsKey($_)) {
        $Fish[$_] = 1
    }
    else {
        $Fish[$_] += 1
    }
}

<# Deal with initial empty numbers #>
foreach ($InitNum in 0..8) {
    if (-not $Fish.Contains("$InitNum")) {
        $Fish["$InitNum"] = 0
    }
}

foreach ($Number in 1..$Days) {

    $FishNextGen = $Fish.Clone()

    $Fish.GetEnumerator() |
    Sort-Object Name -Descending |
    ForEach-Object -Process {
        $NextNumber = "$($_.Name - 1)"


        if ($NextNumber -eq '-1') {
            $FishNextGen['6'] += $Fish['0']
            $FishNextGen['8'] = $Fish['0']
        }
        else {
            $FishNextGen[$NextNumber] = $Fish["$($_.Name)"]
        }
    }

    $Fish = $FishNextGen
}

$Fish.GetEnumerator() |
ForEach-Object -Begin {
    [long] $FinalCount = 0
} -Process {
    $FinalCount += $_.Value
}

@'
Final Count: {0}
'@ -f $FinalCount
