BeforeAll -Scriptblock {
    . $PSScriptRoot\part02.ps1
}

Describe -Name 'part 02' -Fixture {
    Context -Name 'Sample Input' -Fixture {

        BeforeAll -Scriptblock {
            $GroupAnswers = Get-AOCGroupAnswersProper -Path $PSScriptRoot\input_sample.txt
        }

        It -Name 'returns <total> for group <group>' -TestCases @(
            @{ total = 3; group = 0 }
            @{ total = 0; group = 1 }
            @{ total = 1; group = 2 }
            @{ total = 1; group = 3 }
            @{ total = 1; group = 4 }
        ) -Test {
            param (
                $total,
                $group 
            )
            
            $GroupAnswers[$group].AllYes | Should -Be $total
        }
    }
}