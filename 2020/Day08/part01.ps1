function Get-AOCAccValue {
    [CmdletBinding()]
    param (
        # Input path
        [Parameter(Mandatory)]
        [string] $Path
    )
    
    begin {

        $valueCounter = 0
        $position = 0
        $lastChange = 0

    }
    
    process {
        
        $contents = (Get-Content -Path $Path -Raw) -split '(\r?\n)' |
            Where-Object { -not [string]::IsNullOrWhiteSpace($_) }
        
        $instructions = $contents | ForEach-Object -Process {

            $toDo, $thisMuch = $_ -split ' '

            [PSCustomObject]@{
                Index = $index++
                Full = $_
                Action = $toDo.Trim()
                Amount = $thisMuch
                Hits = 0
            }

        }

        $seen = [System.Collections.Generic.SortedSet[int]]::new()
        1..1000 | ForEach-Object -Begin {
            $run = $true
        } -Process {
            $counter = $_
            if ($seen.Contains($position) -and $run) {
                $run = $false
                Write-PSFMessage -Level Warning -Message "Duplicate Index: $position"
            }

            if ($run) {
                $seen.add($position) | Out-Null

                $ci = $instructions | Where-Object Index -eq $position

                switch ($ci.Action) {
                    'nop' {
                        $ci.Hits++
                        $position++
                        break
                    }
                    'acc' {
                        $ci.Hits++
                        $position++
                        $valueCounter += [int]($ci.Amount)
                        break
                    }
                    'jmp' {
                        $ci.Hits++
                        $position += [int]($ci.Amount)
                        break
                    } default {
                        $run = $false
                        break
                    }
                }

                Write-PSFMessage "Index $($ci.Index) - Line: $($ci.Full) - Hits: $($ci.Hits) - Value: $valueCounter"
                $ci | Select-Object @{
                        Name = 'ID'
                        Expression = { $counter }
                    },
                    *,
                    @{
                        Name = 'FinalValue'
                        Expression = { $valueCounter }
                    }, @{
                        Name = 'LastChange'
                        Expression = { [int]($ci.Amount) }
                    },
                    @{
                        Name = 'BeforeChange'
                        Expression = { $valueCounter + [int]($ci.Amount) }
                    }
                

            }
        }

    }
    
    end {
        
    }
}