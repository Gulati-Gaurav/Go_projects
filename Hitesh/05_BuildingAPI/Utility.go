package main

// helper functions / middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
