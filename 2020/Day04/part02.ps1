<#

    byr (Birth Year)
    iyr (Issue Year)
    eyr (Expiration Year)
    hgt (Height)
    hcl (Hair Color)
    ecl (Eye Color)
    pid (Passport ID)
    cid (Country ID)

    byr valid:   2002
    byr invalid: 2003

    hgt valid:   60in
    hgt valid:   190cm
    hgt invalid: 190in
    hgt invalid: 190

    hcl valid:   #123abc
    hcl invalid: #123abz
    hcl invalid: 123abc

    ecl valid:   brn
    ecl invalid: wat

    pid valid:   000000001
    pid invalid: 0123456789
#>
function ConvertFrom-AOCPassword {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory)]
        [string]$Path
    )

    process {
        $Content = Get-Content -Path $Path

        $Content | ForEach-Object -Begin { $index = 1; $hash = @{} } -Process {
            if ($PSItem -match "^\s*$") {
                $index++
            }

            $hash[$index] = $hash[$index] + ' ' + $PSItem
        }

        foreach ($passport in $hash.GetEnumerator()) {
            Write-Verbose -Message "`nTesting: $($passport.Value)"

            $regexAll = '^(?=.*\bbyr\b:(?<byr>.+?\b))(?=.*\biyr\b:(?<iyr>.+?\b))(?=.*\beyr\b:(?<eyr>.+?\b))(?=.*\bhgt\b:(?<hgt>.+?\b))(?=.*\bhcl\b:(?<hcl>#?.{6})\b)(?=.*\becl\b:(?<ecl>.+?\b))(?=.*\bpid\b:(?<pid>.{9})\b).*$'
            $matchesRegex = $passport.Value -match $regexAll

            $potentialPassport = [PSCustomObject]@{
                Line = $passport.Value
                byr = $Matches['byr'] 
                hgt = $Matches['hgt']
                hcl = $Matches['hcl']
                ecl = $Matches['ecl']
                _pid = $Matches['pid']
                iyr = $Matches['iyr']
                eyr = $Matches['eyr']
            }

            $isValid = $true
            $message = $null

            $byr = $potentialPassport.byr
            if (-not ($byr -ge 1920 -and $byr -le 2002)) {
                $Message += "byr invalid: $byr | $($potentialPassport.byr) <> 1920 - 2002"
                $isValid = $false
            }

            $iyr = $potentialPassport.iyr -as [int]
            if (-not ($iyr -ge 2010 -and $iyr -le 2020)) {
                $Message += "iyr invalid: $iyr | $($potentialPassport.iyr) <> 2010 - 2020"
                $isValid = $false
            }

            $eyr = $potentialPassport.eyr -as [int]
            if (-not ($eyr -ge 2020 -and $eyr -le 2030)) {
                $Message += "eyr invalid: $eyr | $($potentialPassport.eyr) <> 2020 - 2030"
                $isValid = $false
            }

            $hgt = $potentialPassport.hgt
            if ($hgt -notmatch '(?<val>\d{3})(?<mea>cm)|(?<val>\d{2})(?<mea>in)') {
                $Message += "hgt invalid: $hgt | $($potentialPassport.hgt)"
                $isValid = $false
            }
            if ($Matches['mea'] -eq 'in' -and ($Matches['val'] -lt 59 -or $Matches['val'] -gt 76)) {
                $Message += "hgt cm invalid: $hgt <> 59-76"
                $isValid = $false
            }
            if ($Matches['mea'] -eq 'cm' -and ($Matches['val'] -lt 150 -or $Matches['val'] -gt 193)) {
                $Message += "hgt in invalid: $hgt <> 150-193"
                $isValid = $false
            }

            $hcl = $potentialPassport.hcl
            if ($hcl -notmatch '#[0-9a-f]{6}') {
                $Message += "hcl invalid: $hcl | $($potentialPassport.hcl)"
                $isValid = $false
            }

            $ecl = $potentialPassport.ecl
            if ($ecl -notin 'amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth') {
                $Message += "ecl invalid: $ecl | $($potentialPassport.ecl)"
                $isValid = $false
            }

            $_pid = $potentialPassport._pid
            if ($_pid -notmatch '\d{9}') {
                $Message += "_pid invalid: $_pid | $($potentialPassport._pid)"
                $isValid = $false
            }

            $potentialPassport |
                Select-Object -Property *, @{
                    Name = 'IsValid'
                    Expression = { $IsValid }
                }, @{
                    Name = 'Errors'
                    Expression = { $Message }
                }
        }
    }
}
