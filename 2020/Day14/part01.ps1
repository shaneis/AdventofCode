[CmdletBinding()]
param (
    [Parameter()]
    [string]
    $Path
)

$InputContent = Get-Content -Path $Path

<#
    $Memory, $Value = $InputContent -split ' = '
    $Memory
    $Value
#>
$Program = @{}

foreach ($Instruction in $InputContent) {
    $Memory, $Value = $Instruction -split ' = '

    if ($Memory -eq 'mask') {
        $Mask = $value
        continue
    }

    $Location = $Memory -match ('mem\[(?<local>\d+)\]') | Select-Object -Property @{ Name = 'Local'; Expression = { $Matches['local'] }}

    $ValueDecimal = [Convert]::ToString($Value, 2).ToCharArray()
    
    $char_string = [Text.StringBuilder]::new()
    for ($i = -1; $i -ge (0 - $Mask.Length); $i--) {
        $char = switch ($Mask[$i]) {
            '0' { '0' }
            '1' { '1' }
            'X' { if ($null -ne $ValueDecimal[$i]) { $ValueDecimal[$i] } else { '0' }}
        }
        $char_string.Insert(0, $char) | Out-Null
    }
    $Program[$Location.Local] = [Convert]::ToInt64($char_string, 2)
}

$Program.GetEnumerator() | Measure-Object -Sum -Property Value