package main

import (
	"GoReader/models"

)

type SideBarLibraries struct {
	pinned []models.Library
	recent []models.Library
}

func(a *App) getLibrariesForSideBar() SideBarLibraries{
	var pinnedLibrary []models.Library;
	a.DB.Model(&models.Library{}).Where("pin_status = ?", true).Find(&pinnedLibrary)

	var recent []models.Library;
	a.DB.Model(&models.Library{}).Where("pin_status = ?", false).Order("last_accessed").Limit(5).Find(&recent)

	return SideBarLibraries{
		pinned: pinnedLibrary,
		recent: recent,
	}

}

