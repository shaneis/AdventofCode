<#

    byr (Birth Year)
    iyr (Issue Year)
    eyr (Expiration Year)
    hgt (Height)
    hcl (Hair Color)
    ecl (Eye Color)
    pid (Passport ID)
    cid (Country ID)

#>

[CmdletBinding()]
param (
    [switch]$UseSampleInput
)

$Content = if ($UseSampleInput) {
    Get-Content -Path "$($PSScriptRoot)\Day04_part01_Sample.txt"
} else {
    Get-Content -Path "$($PSScriptRoot)\Day04_part01.txt"
}

$Content | ForEach-Object -Begin { $index = 1; $hash = @{} } -Process {
    if ($PSItem -match "^\s*$") {
        $index++
    }

    $hash[$index] = $hash[$index] + ' ' + $PSItem
}

$hash.GetEnumerator() | ForEach-Object -Begin { $ValidPassports = 0 } -Process {
    if ($PSItem.Value -match '^(?=.*\bbyr\b)(?=.*\biyr\b)(?=.*\beyr\b)(?=.*\bhgt\b)(?=.*\bhcl\b)(?=.*\becl\b)(?=.*\bpid\b).*$') {
        $ValidPassports++
    }
}

$ValidPassports