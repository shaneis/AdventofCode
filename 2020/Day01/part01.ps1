#[int[]]$Sample = 1721, 979, 366, 299, 675, 1456
[int[]]$Sample = Get-Content -Path .\day01_part01.txt

$list = [System.Collections.Generic.List[int]]::new()
$list.AddRange($Sample)


foreach ($number in $list.ToArray()) {
    foreach ($otherNumber in $list.ToArray()) {
        if (($number + $otherNumber) -eq 2020) {
            "$number * $otherNumber = $($number * $otherNumber)"
            return
        }
    }
    "Nope! Removing: $number - $($list.Count) remaining..."
    $list.Remove($number) | Out-Null
}

