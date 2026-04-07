// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package main

import (
	"context"
	"fmt"
	"log"

	docspace_api_sdk "github.com/ONLYOFFICE/docspace-api-sdk-go/v3"
)

func main() {
	cfg := docspace_api_sdk.NewConfiguration()
	cfg.Servers = docspace_api_sdk.ServerConfigurations{
		{
			URL:         "http://localhost:8092",
			Description: "Local DocSpace",
		},
	}

	client := docspace_api_sdk.NewAPIClient(cfg)
	ctx := context.Background()

	authRequest := docspace_api_sdk.NewAuthRequestsDtoWithDefaults()
	authRequest.SetUserName("example@onlyoffice.com")
	authRequest.SetPassword("11111111")

	authResponse, _, err := client.AuthenticationAPI.
		AuthenticateMe(ctx).
		AuthRequestsDto(*authRequest).
		Execute()
	if err != nil {
		log.Fatalf("authentication failed: %v", err)
	}

	authPayload := authResponse.GetResponse()
	token := authPayload.GetToken()
	ctx = context.WithValue(ctx, docspace_api_sdk.ContextAccessToken, token)

	myFolder, _, err := client.FilesFoldersAPI.GetMyFolder(ctx).Execute()
	if err != nil {
		log.Fatalf("failed to get my folder: %v", err)
	}

	myFolderPayload := myFolder.GetResponse()
	currentFolder := myFolderPayload.GetCurrent()
	parentFolderID := currentFolder.GetId()
	folderPayload := docspace_api_sdk.NewCreateFolderWithDefaults()
	folderPayload.SetTitle("SDK Sample Folder")

	createdFolder, _, err := client.FilesFoldersAPI.
		CreateFolder(ctx, parentFolderID).
		CreateFolder(*folderPayload).
		Execute()
	if err != nil {
		log.Fatalf("failed to create folder: %v", err)
	}

	createdFolderPayload := createdFolder.GetResponse()
	fmt.Printf("Created folder ID: %d\n", createdFolderPayload.GetId())
}
