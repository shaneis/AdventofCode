BeforeAll -Scriptblock {
    . $PSScriptRoot\part01.ps1
}

Describe -Name 'part 01' -Fixture {
    Context -Name 'Sample Input' -Fixture {
        It -Name 'returns expected result <Number> for colour <Colour>' -TestCases @(
            @{ Colour = 'shiny gold'; Number = 4}
        ) -Test {
            param(
                $Colour,
                $Number
            )
            $result = Get-AOCBagHolders -Path $PSScriptRoot\input_sample.txt -Colour $Colour
            $potentialBags = $result | Select-Object -Property Bag -Unique | Measure-Object

            $potentialBags.Count | Should -Be $Number
        }
    }
}