function Get-AOCGroupAnswersProper {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory)]
        [string]$Path
    )
    
    begin {

    }
    
    process {
        $groups = (Get-Content -Path $Path -Raw) -split '(\r?\n){2}' | Where-Object { -not [String]::IsNullOrWhitespace($PSItem) }
        
        foreach ($group in $groups) {

            $individuals = $group -split '\r?\n' | Where-Object { -not [String]::IsNullOrWhitespace($PSItem) }

            $allYes = $individuals.ToCharArray() |
                Group-Object |
                Where-Object Count -eq $individuals.Count |
                Measure-Object

            Write-Verbose "$($allYes.Count) for $($individuals.count)"

            [PSCustomObject]@{
                Counter = ++$index
                GroupLine = $group -replace '\r?\n', ' '
                Individuals = $individuals -join ', '
                Choices = $individuals.ToCharArray() -join ', '
                AllYes = $allYes.Count
            }
        }
    }
    
    end {
        
    }
}