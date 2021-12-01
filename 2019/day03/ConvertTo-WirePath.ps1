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