function Get-AOCAccValue_Fix {
    [CmdletBinding()]
    param (
        # Input path
        [Parameter(Mandatory)]
        [string] $Path
    )
    
    begin {

        $valueCounter = 0
        $position = 0
        $oldPos = $position

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
        $run = $true
        foreach ($counter in 1..1000 ) {

            if ($counter -eq 1000) {
                Write-PSFMessage -Level Critical -Message "Increase runs!"
            }

            if ($seen.Contains($position) -and $run) {
                $run = $false
                #Write-PSFMessage -Level Verbose -Message "Duplicate Index: $position"
                return
            }

            if ($run) {
                $seen.add($position) | Out-Null

                $ci = $instructions | Where-Object Index -eq $position

                switch ($ci.Action) {
                    'nop' {
                        $ci.Hits++
                        $oldPos = $position
                        $position++
                        break
                    }
                    'acc' {
                        $ci.Hits++
                        $oldPos = $position
                        $position++
                        $valueCounter += [int]($ci.Amount)
                        break
                    }
                    'jmp' {
                        $ci.Hits++
                        $oldPos = $position
                        $position += [int]($ci.Amount)
                        break
                    } default {
                        $oldPos = $position
                        $position = $contents.Count + 1
                        $run = $false
                        break
                    }
                }

                $ci | Select-Object @{
                        Name = 'ID'
                        Expression = { $counter }
                    },
                    *,
                    @{
                        Name = 'FinalValue'
                        Expression = { $valueCounter }
                    },
                    @{
                        Name = 'WasSuccessful'
                        Expression = { if ($position -ge $contents.Count) { $true } else { $false } }
                    }, @{
                        Name = 'Positions'; Expression = { " $position | $oldPos" }
                    }, @{
                        Name = 'ConCount'; Expression = { $contents.Count }
                    }

                if ($position -ge $contents.count -or $position -eq $oldPos) { return }
            }
        }

    }
    
    end {
        
    }
}

