function Get-AOCPreAmble {
    [CmdletBinding()]
    param (
        $path,
        $preamble_end = 24
    )
    
    begin {
        $start = 0
        $found = $false
    }
    
    process {
        $input_contents = Get-Content -Path $path

        $end = $preamble_end
        $rolling_preamble = $input_contents[$start..$end]

        $test_preamble = $rolling_preamble.Clone()
        
        foreach ($iter in (0..$input_contents.Count)) {

            $next_number = $input_contents[$end+1]

            $does_match = $false
            foreach ($t in $test_preamble) {
                $test_number = $next_number - $t
                if ($test_number -in $test_preamble -and $test_number -ne $t) {
                    Write-PSFMessage -Level Verbose -Message "SubSet! $test_number"
                    $does_match = $true
                    break
                }
            }

            if (-not $does_match) {
                [PSCustomObject]@{
                    NoMatch = $next_number
                }
                return
            }

            $start +=1
            $end += 1
            $test_preamble = $input_contents[$start..$end]
        }

        
    }
    
    end {
        
    }
}