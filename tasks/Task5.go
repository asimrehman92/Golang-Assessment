package tasks

import "github.com/apex/log"

func Task5() {
	log.Info("Here is Task 5")

	q := `WITH RankedEmployees AS (
		SELECT 
		  e.name AS Employee, 
		  e.salary AS Salary, 
		  d.name AS Department,
		  ROW_NUMBER() OVER(PARTITION BY e.departmentId ORDER BY e.salary DESC) AS Rank
		FROM 
		  Employee e
		JOIN 
		  Department d ON e.departmentId = d.id
	  )
	  SELECT 
		Department, 
		Employee, 
		Salary
	  FROM 
		RankedEmployees
	  WHERE 
		Rank <= 3;`
	log.Debug(q)
}
