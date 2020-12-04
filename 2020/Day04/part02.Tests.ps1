BeforeAll -ScriptBlock {
    . "$($PSScriptRoot)\part02.ps1"
}

Describe -Name "ConvertFrom-AOCPassword" -Fixture {
    Context -Name "returns failure passwords correctly" -Fixture {
        It -Name 'marks failures correctly' -Test {
            $results = ConvertFrom-AOCPassword -Path "$($PSScriptRoot)\test02_fail.txt"
            ($results | Select-Object -Property IsValid -Unique).IsValid | Should -BeFalse
            $results.Count | Should -Be 4
        }
    }

    Context -Name 'returns success correctly' -Fixture {
        It -Name 'marks failures correctly' -Test {
            $results = ConvertFrom-AOCPassword -Path "$($PSScriptRoot)\test02_pass.txt"
            ($results | Select-Object -Property IsValid -Unique).IsValid | Should -BeTrue
            $results.Count | Should -Be 4
        }
    }
}