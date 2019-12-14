$here = Split-Path -Parent $MyInvocation.MyCommand.Path
$sut = (Split-Path -Leaf $MyInvocation.MyCommand.Path) -replace '\.Tests\.', '.'
. "$here\$sut"

Describe "02_part01" {
    It -Name 'should have a mandatory parameter called <Parameter>' -TestCases @(
        @{Parameter = 'ProgramInput'}
    ) -Test {
        param (
            $Parameter
        )

        $Function = Get-Command -Name Resolve-GravityAssistProgram
        $Function | Should -HaveParameter $Parameter -Mandatory
    }
    
    It "should return expected output <Result> with given input <ProgramInput>" -TestCases @(
        @{ProgramInput = 1,0,0,0,99; Result = 2,0,0,0,99}
    ) {
        param (
            $ProgramInput,
            $Result
        )
        $TestRun = Resolve-GravityAssistProgram -ProgramInput $ProgramInput
        $TestRun.OutputProgram | Should -BeExactly $Result
    }
}
