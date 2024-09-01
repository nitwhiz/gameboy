package addr

const (
	LCDC_ENABLE                 = byte(7)
	LCDC_WINDOW_TILE_MAP_SELECT = byte(6)
	LCDC_WINDOW_DISPLAY_ENABLE  = byte(5)
	LCDC_TILE_DATA_SELECT       = byte(4)
	LCDC_BG_TILE_MAP_SELECT     = byte(3)
	LCDC_SPRITE_SIZE            = byte(2)
	LCDC_SPRITE_ENABLE          = byte(1)
	LCDC_BG_WINDOW_ENABLE       = byte(0)
)

const (
	STAT_LYCLY_INTERRUPT_ENABLE    = byte(6)
	STAT_OAM_SCAN_INTERRUPT_ENABLE = byte(5)
	STAT_VBLANK_INTERRUPT_ENABLE   = byte(4)
	STAT_HBLANK_INTERRUPT_ENABLE   = byte(3)
	STAT_COINCIDENCE_FLAG          = byte(2)
)

const (
	// OAM_ATTR_PRIORITY 0 = No, 1 = BG and Window colors 1–3 are drawn over this OBJ
	OAM_ATTR_PRIORITY = byte(7)
	OAM_ATTR_Y_FLIP   = byte(6)
	OAM_ATTR_X_FLIP   = byte(5)
	OAM_ATTR_PALETTE  = byte(4)
)
