[CmdletBinding()]
param (
    $path,
    $search = 41682220 # my answer from part 1
)
measure-script {
    #$nums = Get-Content -path $path
$nums = Get-Content -path .\input.txt
$search = 41682220

<#

foreach ($counter in (0..$numCount)) {
    $nums[(($numCount) - ($numCount - $counter))..($numCount - 1)] | Get-ContiguousRange -Search $search
}
#>

$n = [Collections.Generic.List[Int64]]::new()
foreach ($i in $nums) {
    $n.Add($i) | Out-Null
}
$end = $nums.Count
$j = 1

foreach ($size in (0..($n.Count -1))) {
    $prevValue = [Int64]::MinValue
    Write-PSFMessage -Level Verbose "Size $size out of $($n.Count -1)"
    foreach ($iter in 1..$n.Count) {
        if ($n[0] -le $search -or $iter -gt 2) {
            #Write-PSFMessage -Level Verbose "Iter $iter out of $($n.Count)"
            $meas = $n | Select-Object -First $iter -OutVariable cont

            if ($meas | Where-Object { $_ -gt $search }) {
                #Write-PSFMessage -level Verbose -Message "$iter - Skipping, items are over"
                break
            }

            $meas = $meas | Measure-Object -Sum 

            if ($prevValue -gt $search) {
                #Write-PSFMessage -level Verbose -Message "$iter - Skipping, sum was already over"
                $prevValue = $meas.Sum
                break
            }

            if ($meas.Sum -eq $search -and $meas.Count -ne 1) {
                $meas | ForEach-Object -Process {
                    [PSCustomObject]@{
                        List = $n
                        Cont = $cont
                        Smallest = $cont | Sort-Object | Select-Object -First 1
                        Largest = $cont | Sort-Object | Select-Object -Last 1
                        Sum = $_.Sum
                        Count = $_.Count
                    }
                    return
                }
            }
        }
    }
    $n = $n[$j..$end]
}
}