package model

type Upgrades struct {
	Class        string `json:"class"`
	UpgradeLevel int64  `json:"upgradeLevel"`
	Cost         int64  `json:"cost"`
}

/*○ Engine
■ Affects Power
■ Upgrade increases Power
○ Turbo
■ Affects Power
■ Upgrade increases Power
○ Intake
■ Affects Power
■ Upgrade increases Power
○ Nitrous
■ Increases Power for a certain period of time
■ Upgrade effects both:
● Amount of Power increased, and
● Period of time for which Nitrous would be enabled for
○ Period of time would be in seconds
○ Body
■ Affects Grip and Weight
■ Upgrade increases Grip and decreases Weight
○ Tires
■ Affects Grip
■ Upgrade Increases Grip
○ Transmission
■ Affects Shift-Time */
