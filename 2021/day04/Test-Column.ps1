#! /usr/bin/pwsh

function Test-Column {
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
		[int] $Column,

		[Parameter(
			Mandatory,
			ValueFromPipelineByPropertyName
		)]
		[int[]] $CalledNumbers
	)

	begin {
		$i = 0
		$IsComplete = $true
	}

	process {
		$Columns = foreach ($l in $Board -split '\r?\n') {
			[int[]] $Numbers = $l -split ' ' | Where-Object {$_ -as [int] -or $_ -eq '0'}
			foreach ($n in $Numbers) {
				Write-Verbose "Checking: $n - Column: $(($i % 5) +1)"
				if ((($i % 5) + 1) -eq $Column) {
					[PSCustomObject] @{
						ExpectedColumn = $Column
						ActualColumn = ($i % 5) + 1
						Number = $n
						CalledNumbers = $CalledNumbers -join ','
					}
				}
				$i++
			}
		}
		
		foreach ($n in $Columns.Number) {
			if ($n -notin $CalledNumbers) {
				$IsComplete = $false
				break
			}
		}	

		[PSCustomObject] @{
			BoardColumn = $Columns.Number -join ' '
			Column = $Column
			CalledNumbers = $Columns[0].CalledNumbers
			IsComplete = $IsComplete
		}
	}
}

