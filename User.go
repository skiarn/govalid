package govalid
import (
  "regexp"
)

func Email(email string) bool{
  Reg := regexp.MustCompile(`^[a-z0-9._&+\*?-]+@[a-z0-9.\-]+\.[a-zA-Z]{2,7}$`)
  return if(len(email) > 5 && Reg.MatchString(email)
}
