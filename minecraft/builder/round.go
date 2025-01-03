package builder

import "phoenixbuilder/minecraft/mctype"

func Round(config *mctype.MainConfig, blc chan *mctype.Module) error {
	Radius := config.Radius
	Facing := config.Facing
	point := config.Position
	switch Facing {
	case "x":
		for i := -Radius; i <= Radius; i++ {
			for j := -Radius; j <= Radius; j++ {
				if i*i+j*j < Radius*Radius {
					var b mctype.Module
					b.Point = mctype.Position{X: point.X, Y: point.Y + i, Z: point.Z + j}
					blc <- &b
				}
			}
		}
	case "y":
		for i := -Radius; i <= Radius; i++ {
			for j := -Radius; j <= Radius; j++ {
				if i*i+j*j < Radius*Radius {
					var b mctype.Module
					b.Point = mctype.Position{X: point.X + i, Y: point.Y, Z: point.Z + j}
					blc <- &b
				}
			}
		}
	case "z":
		for i := -Radius; i <= Radius; i++ {
			for j := -Radius; j <= Radius; j++ {
				if i*i+j*j < Radius*Radius {
					var b mctype.Module
					b.Point = mctype.Position{point.X + i, point.Y + j, point.Z}
					blc <- &b
				}
			}
		}
	}
	return nil
}
