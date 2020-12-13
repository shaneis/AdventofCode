

[int[]]$c = Get-Content -Path $pwd\input.txt

$highestRatedAdapter = $c | Measure-Object -Maximum | ForEach-Object { $_.Maximum + 3 }

#$c

#$highestRatedAdapter

$prev = 0
$chain = foreach ( $i in $c | Sort-Object) {

    [PSCustomObject]@{
        Adapter = $i
        Prev = $prev
        Difference = $i - $prev
    }

    $prev = $i
}

$chain += [PSCustomObject]@{
    Adapter = $highestRatedAdapter
    Prev = $chain[-1].Adapter
    Difference = $highestRatedAdapter - ($chain[-1].Adapter)
}

$Answer = $chain | Group-Object Difference -NoElement

($Answer.GetValue(0).Count) * ($Answer.GetValue(1).Count)