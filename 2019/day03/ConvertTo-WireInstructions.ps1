function ConvertTo-WireInstructions {
	[CmdletBinding()]
	param(
		[Parameter(
			Mandatory,
			ValueFromPipeline,
			ValueFromPipelineByPropertyName
		)]
		[string] $Wire
	)

	process {
		foreach ($WireInstruction in $Wire) {
			Write-Verbose "Working on: [$WireInstruction]"
			$Instructions = $WireInstruction -split ','

			$Iteration = 1

			foreach ($Instruction in $Instructions) {
				Write-Verbose "Working on instruction: [$Instruction]"

				$XDelta = 0
				$YDelta = 0

				switch ($Instruction) {
					{$Instruction.Substring(0, 1) -eq 'R'} {$XDelta = 1 * ($Instruction.Substring(1) -as [int])}
					{$Instruction.Substring(0, 1) -eq 'L'} {$XDelta = -1 * ($Instruction.Substring(1) -as [int])}
					{$Instruction.Substring(0, 1) -eq 'U'} {$YDelta = 1 * ($Instruction.Substring(1) -as [int])}
					{$Instruction.Substring(0, 1) -eq 'D'} {$YDelta = -1 * ($Instruction.Substring(1) -as [int])}
				}

				[PSCustomObject] @{
					MoveId = $Iteration
					Instruction = $Instruction
					XDelta = $XDelta
					YDelta = $YDelta
				}

				$Iteration++
			}
		}
	}
}