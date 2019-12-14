function Resolve-GravityAssistProgram {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory,
                   ValueFromPipeline,
                   ValueFromPipelineByPropertyName)]
        [Int[]]
        $ProgramInput
    )
}
