BeforeAll -Scriptblock {
    . $PSScriptRoot\part01.ps1
}

Describe -Name 'part 01' -Fixture {
    Context -Name 'Sample Input' -Fixture {
        It -Name 'returns expected result' -Test {
            $result = Get-AOCGroupAnswers -Path $PSScriptRoot\input_sample.txt
            $AllYes = $result | Measure-Object -Sum -Property AllYes

            $AllYes.Sum | Should -Be 11
        }
    }
}