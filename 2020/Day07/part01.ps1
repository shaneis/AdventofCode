function Get-AOCBagHolders {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory)]
        [string]$Path,

        [Parameter(Mandatory)]
        [string]$Colour,

        [Parameter()]
        [ValidateNotNullOrEmpty()]
        [int]$Count = 1
    )
    
    begin {
        $regexParser = '(?<nu>\d+)? (?<bag_holder>\w+ \w+) (?=bag(s)?)'
    }
    
    process {
        $content = (Get-Content -Path $Path) -split '(\r?\n)' | Where-Object { -not [String]::IsNullOrWhiteSpace($PSItem) }

        $bagMaps = foreach ($rule in $content) {
            $bag = $rule -split 'bag(s)?' | Select-Object -First 1
            
            foreach ($holder in $rule -split ',') {
                $holder -match $regexParser | Out-Null

                [PSCustomObject]@{
                    Bag = $bag.Trim()
                    Number = if ($Matches['bag_holder'] -match 'no other') { 0 } else { $Matches['nu'] }
                    BagHolder = if ($Matches['bag_holder'] -match 'no other') { $null } else { $Matches['bag_holder'] }
                    Count = $count
                    PSTypeName = 'RecursiveAOC'
                }
            }
        }

        Write-Verbose "$($bagMaps | out-String)"
        Get-AOCContainerBags -Map $bagMaps -Bag $Colour
        
    }
    
    end {
        $Seen
    }
}

function Get-AOCContainerBags {
    [CmdletBinding()]
    param (
        # Object to check
        [Parameter(Mandatory)]
        [PSTypeName('RecursiveAOC')]$Map,

        # What to search for
        [Parameter(Mandatory)]
        [string]$Bag
    )
    
    begin {
        $seen = [System.Collections.Generic.SortedSet[string]]::new()
    }
    
    process {
        $Map | Where-Object BagHolder -eq $Bag | ForEach-Object -Process {
            if (-not ($seen.Contains($_.Bag))) {
                Write-Verbose "adding $($_.Bag)"
                $seen.Add($_.Bag) | Out-Null
                Get-AOCContainerBags -Map $Map -Bag $_.Bag
            }
            $_
        }
    }
    
    end {
        
    }
}