package ninja_world_errors

import "fmt"

type ninjaWorldErrors struct {
	err error
	msg string
}

func (err ninjaWorldErrors) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%s: %v", err.msg, err.err)
	}
	return err.msg
}

var (
	VILLAGEALREADYEXISTS       = ninjaWorldErrors{msg: "village already in the world"}
	VILLAGEDOESNOTEXISTS       = ninjaWorldErrors{msg: "village doesn't exists in the world"}
	VILLAGEDESTROYED           = ninjaWorldErrors{msg: "village is in destroyed state"}
	VILLAGEALREADYINLINK       = ninjaWorldErrors{msg: "another village is added in this direction"}
	OTSUTSUKINAMEALREADYEXISTS = ninjaWorldErrors{msg: "Duplicate otsutsuki names found"}
)
