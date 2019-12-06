$here = Split-Path -Parent $MyInvocation.MyCommand.Path
$sut = (Split-Path -Leaf $MyInvocation.MyCommand.Path) -replace '\.Tests\.', '.'
. "$here\$sut"

Describe "01_part01" {
    It "passes the examples given in the description" -TestCases @(
        @{Mass = 12; Result = 2}
        @{Mass = 14; Result = 2}
        @{Mass = 1969; Result = 654}
        @{Mass = 100756; Result = 33583}
    ) -Test {
        param (
            $Mass,
            $Result
        )
        $FuelRequired = Get-FuelPerModuleMass -Mass $Mass
        $FuelRequired.Fuel | Should -Be $Result
    }
}
