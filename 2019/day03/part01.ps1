[CmdletBinding()]
param (
	[string] $Path
)

$TestLines = Get-Content -Path $Path

$Wire01, $Wire02 = $TestLines -split '\r?\n'

Write-Verbose "Wire 01: [$Wire01]"
Write-Verbose "Wire 02: [$Wire02]"

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

function ConvertTo-WirePath {
	[CmdletBinding()]
	param(
		[Parameter(
			Mandatory,
			ValueFromPipeline
		)]
		[psobject] $WireInstruction
	)

	begin {
		$Origin = [PSCustomObject]@{
			X = 0
			Y = 0
		}
	}

	end {
		foreach ($wi in $WireInstruction) {
			Write-Verbose "Working on: [$wi]"
			foreach ($Position in $wi) {
				Write-Verbose "Working on position: [$Position]"
				foreach ($X in 0..$Position.XDelta) {
					$IncrementX = if ($X -eq 0) {0} elseif ($X -lt 0) {-1} else {1}

					$NewOrigin = [PSCustomObject] @{
						X = $Origin.X + $IncrementX
						Y = $Origin.Y
					}
					$NewOrigin

					$Origin = $NewOrigin
				}

				foreach ($Y in 0..$Position.YDelta) {
					$IncrementY = if ($Y -eq 0) {0} elseif ($Y -lt 0) {-1} else {1}

					$NewOrigin = [PSCustomObject]@{
						X = $Origin.X
						Y = $Origin.Y + $IncrementY
					}
					$NewOrigin

					$Origin = $NewOrigin
				}
			}
		}
	}
}

$Wire01Instructions = ConvertTo-WireInstructions -Wire $Wire01
$Wire02Instructions = ConvertTo-WireInstructions -Wire $Wire02

$Wire01Path = ConvertTo-WirePath -WireInstruction $Wire01Instructions
$Wire02Path = ConvertTo-WirePath -WireInstruction $Wire02Instructions

$Results = foreach ($Instr in $Wire01Path) {
	Write-Verbose "Going through: [$Instr]"
	foreach ($SecondPath in $Wire02Path) {
		if ($Instr.X -eq 0 -and $Instr.Y -eq 0) {continue}

		if ($Instr.X -eq $SecondPath.X -and $Instr.Y -eq $SecondPath.Y) {

			$ManhattanDistance = [Math]::Abs($SecondPath.X) + [Math]::Abs($SecondPath.Y)

			[PSCustomObject]@{
				X = $SecondPath.X
				Y = $SecondPath.Y
				ManhattanDistance = $ManhattanDistance
			}
		}
	}
}

$Results | Sort-Object -Property ManhattanDistance
