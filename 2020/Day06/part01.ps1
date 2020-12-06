function Get-AOCGroupAnswers {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory)]
        [string]$Path
    )
    
    begin {
        
        'a'..'z' | ForEach-Object -Begin { $letterHash = @{} } -Process {
            $letterHash.Add($PSItem, 0)
        }
    }
    
    process {
        $groups = (Get-Content -Path $Path -Raw) -split "\r?\n\r?\n" 
        
        foreach ($group in $groups) {

            $individuals = $group -split "\r?\n"

            $choice = foreach ($individual in $individuals) {
                [PSCustomObject]@{
                    Individual = ++$person
                    choices = $individual.ToCharArray()
                }
            }

            [PSCustomObject]@{
                Counter = ++$index
                GroupLine = $group -replace '\r?\n', ' '
                Individuals = $individuals.Count
                Choices = $individuals.ToCharArray()
                AllYes = ($choice.choices | Select-Object -Unique).Count
            }
        }
    }
    
    end {
        
    }
}