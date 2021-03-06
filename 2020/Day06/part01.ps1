function Get-AOCGroupAnswers {
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

            $choice = foreach ($individual in $individuals) {
                [PSCustomObject]@{
                    Individual = ++$person
                    choices = $individual.ToCharArray()
                }
            }

            [PSCustomObject]@{
                Counter = ++$index
                GroupLine = $group -replace '\r?\n', ' '
                Individuals = $individuals
                Choices = $individuals.ToCharArray()
                AllYes = ($choice.choices | Select-Object -Unique).Count
            }
        }
    }
    
    end {
        
    }
}