#! /usr/bin/pwsh

function Test-Row {
    [CmdletBinding()]
	param (
		[Parameter(
			Mandatory,
			ValueFromPipeline,
			ValueFromPipelineByPropertyName
		)]
		[Alias('BingoBoard')]
		[string] $Board,

		[Parameter(
			Mandatory,
			ValueFromPipelineByPropertyName
		)]
		[int] $Row,

		[Parameter(
			Mandatory,
			ValueFromPipelineByPropertyName
		)]
		[Alias('Numbers')]
		[int[]] $CalledNumbers
	)

	begin {
		$IsComplete = $true
	}

	process {
		foreach ($BoardRow in $Board -split '\r?\n') {
			foreach ($n in ($BoardRow -split ' ')) {
				if ($n -notin $CalledNumbers) {
					$IsComplete = $false
					break
				}
			}

			[PSCustomObject] @{
				BoardRow = $BoardRow
				Row = $Row
				Numbers = $CalledNumbers -join ','
				IsComplete = $IsComplete
			}
		}
	}	
}
