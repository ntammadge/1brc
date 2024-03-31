# My Implementation process

Everything that follows is a log of my thought process while implementing a solution.

## Initial implementation

Getting a functional solution is the first step, but I know it'll need to be more performant that whatever I come up with in this attempt. My initial considerations:
- Reading the entire file (~12GB) into memory prior to processing is going to consume a lot of RAM and doing so entirely before processing any lines won't be performant
- Attempt to minimize the amount of data and calculations done. Example: no need to store every reading from each station for the average calculation because that's extra execution time resizing the collection holding that data, and it's just a `sum of temps / number of readings` calculation, which can be done while handling each line.
- Start simple

I've verified my output matches the output of the baseline implementation, nearly. It appears there's an encoding difference between the baseline Java implementation and my Go implementation. The data looks equivalent for the 415 stations in my `measurements.txt` outside of that. There's one odd station in both outputs that doesn't have any data associated with it, but it's the same oddity in both outputs so I'm not going to think about it too much at the moment. I can generate additional measurement files if I want to try to avoid it.

I ran [calculate_average_baseline.sh](../../../../calculate_average_baseline.sh) to see what kind of output I should be expecting and the runtime. Runtime via `time ./calculate_average_baseline.sh` was 3 minutes 11.293 seconds (real) on my local machine. I saw some discussion about `time` not having great accuracy, but it'll do for now. I can come up with something else when I start caring more (ex. when execution times get short).

I've created [run.sh](./run.sh) so I can at least replicate the process running inside of a script to try to have an even comparison, if it matters. My implementation ran in 2 minutes 46.215s (real) via `time ./run.sh`. Beating the baseline implementation is a good start. I'm not quite sure how much of that is implementation related and executing Java vs executing Go.

I've looked at the baseline implementation in [CalculateAverage_baseline.java](../../java/dev/morling/onebrc/CalculateAverage_baseline.java) to see if there's anything interesting happening there and found the station data appears to be inserted into a red-black tree, which would effectively auto-sort the station data for me. It'd be interesting to see the difference in performance between a map and a sorted tree. The baseline is also using built-in min/max methods, which is something do if Go's implementations are more optimized than simple `if val > high` style checks. Other things I know I'll need to consider:
- Concurrent reading and processing of data
- Is there a faster way to read the file?
- Is there a faster way to parse each line's data?

There's likely more that I'm unaware of, but these are some known unknowns.