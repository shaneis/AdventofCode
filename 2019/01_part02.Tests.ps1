$here = Split-Path -Parent $MyInvocation.MyCommand.Path
$sut = (Split-Path -Leaf $MyInvocation.MyCommand.Path) -replace '\.Tests\.', '.'
. "$here\$sut"

Describe "01_part02" {
    It "passes the examples given in the description" -TestCases @(
        @{Mass = 14; Result = 2}
        @{Mass = 1969; Result = 966}
        @{Mass = 100756; Result = 50346}
    ) -Test {
        param (
            $Mass,
            $Result
        )
        $FuelWithFuel = Get-FuelPerModuleMassWithFuel -Mass $Mass
        $FuelWithFuel.TotalFuel | Should -Be $Result
    }

    It 'works expected <Result> when <Mass> passed in from the pipeline' -TestCases @(
        @{Mass = 14; Result = 2 }
        @{Mass = 1969; Result = 966 }
        @{Mass = 100756; Result = 50346 }
    ) -Test {
        param (
            $Mass,
            $Result
        )
        $FuelWithFuel = $Mass | Get-FuelPerModuleMassWithFuel
        $FuelWithFuel.TotalFuel | Should -Be $Result
    }

    It 'works expected <Result> when <Mass> passed in from the pipeline' -TestCases @(
        @{Mass = 14, 1969; Result = (2, 966) }
    ) -Test {
        param (
            $Mass,
            $Result
        )
        $FuelWithFuel = $Mass | Get-FuelPerModuleMassWithFuel
        $FuelWithFuel.TotalFuel[0] | Should -BeIn $Result
        $FuelWithFuel.TotalFuel[1] | Should -BeIn $Result
    }
}
