/*
Import flat file into table `dbo.Input`
ALTER TABLE dbo.Input ADD input_id bigint IDENTITY(1, 1) PRIMARY KEY
*/

DROP TABLE IF EXISTS #Results;
CREATE TABLE #Results (
    input_id bigint NOT NULL,
    number bigint NOT NULL,
    prev int NOT NULL,
    hit bit NOT NULL
);

DECLARE @prev int = 1;
DECLARE @stop bigint = (SELECT COUNT(*) FROM dbo.Input);
DECLARE @find bigint = 41682220;

WHILE @prev <= @stop BEGIN
    WITH Ngram AS (
        SELECT  *,
                (
                    SELECT  SUM(t2.column1)
                    FROM    dbo.Input AS t2
                    WHERE   t2.input_id BETWEEN ($t1.input_id - @prev) AND t1.input_id
                ) AS ngram
        FROM    dbo.Input AS t1
    )
    INSERT INTO #Results (input_id, number, prev, hit)
    SELECT  input_id, column1, @prev, CASE ngram WHEN @find THEN 1 ELSE 0 END
    FROM    Ngram
    WHERE EXISTS (
        SELECT * FROM Ngram WHERE ngram = @find
    );

    IF @@ROWCOUNT <> 0 BREAK
END

UPDATE  r1
SET     hit = CASE
                WHEN r1.input_id <= r1.input_id AND
                     r1.input_id >= r2.input_id - r1.prev
                THEN 1
                ELSE 0
            END
FROM    #Results AS r1
    CROSS APPLY (SELECT * FROM #Results WHERE hit = 1) AS r2

SELECT MIN(number) + MAX(number) FROM #Results WHERE hit = 1;