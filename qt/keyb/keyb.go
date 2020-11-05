package keyb

import (
	//"fmt"
	"runtime"
	"testing"

	"github.com/therecipe/qt/core"
)

// In a typical Windows keyboard:
// modif1 means Shift key
// modif2 means Control key
// modif3 means Win key
// modif4 means Alt key
var (
	modif1     = int(core.Qt__ShiftModifier)
	modif2     = int(core.Qt__ControlModifier)
	modif3     = int(core.Qt__MetaModifier)
	modif4     = int(core.Qt__AltModifier)
	noModifier = int(core.Qt__NoModifier)
)

// OSKeyModifiers changes modif keys meaning to accomodate to macos systems.
func OSKeyModifiers() (int, int, int, int, int) {
	if runtime.GOOS == "darwin" {
		modif2 = int(core.Qt__MetaModifier)
		modif3 = int(core.Qt__AltModifier)
		modif4 = int(core.Qt__ControlModifier)
	}

	return modif1, modif2, modif3, modif4, modif5
}
