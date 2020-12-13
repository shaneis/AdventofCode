[int[]]$c = Get-Content -Path $pwd\input_sample.txt

[int]$highestRatedAdapter = ($c | Measure-Object -Maximum).Maximum  + 3 

#$full_c = $c + $highestRatedAdapter

$prev = 0

$chain = foreach ( $i in $c | Sort-Object) {

    [PSCustomObject]@{
        Adapter    = $i
        Prev       = $prev
        Difference = $i - $prev
    }

    $prev = $i
}

$chain

$allChains = foreach ($i in (0..($chain.Count -1))) {
    #"A: $($chain[$i])"
    for ($j = $i+1; $j -lt $chain.Count; $j++) {
        #"C: $($chain[$j])"
        if (($chain[$j].Adapter - $chain[$i].Adapter) -le 3 -and
            ($chain[$j].Adapter - $chain[$i].Adapter) -gt 0) {
            [PSCustomObject]@{
                Adapter = $chain[$i].Adapter
                NewA = $chain[$j].Adapter
                Diff = $chain[$i].Difference
                NewDiff = $chain[$j].Adapter - $chain[$i].Adapter
            }
            continue <# potential #>
        }
        else { break <# too big #>}
    }
}


$groupedChain = $allChains | Group-Object -Property Adapter
for ($i = 0; $i -lt $allChains.Count; $i++) {
    $thisgroup = $groupedChain[$i]
    
    $perms = get-new -ThisGroup $thisgroup -Count $thisgroup.Count

    $thisgroup | Where-object { $_.Group}
}
function get-new {
    [CmdletBinding()]
    param (
        [Parameter()]
        $ThisGroup,

        [Parameter()]
        $Count
    )

    process {
        for ($i = 0; $i -lt $Count; $i++) {
            [PSCustomObject]@{
                Full = $ThisGroup.Group[$i]
                Current = $ThisGroup.Group[$i].Adapter
                Next = $ThisGroup.Group[$i].NewA
            }
        }
    }
}