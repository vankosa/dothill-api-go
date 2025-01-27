// Copyright (c) 2021 Enix, SAS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Authors:
// Paul Laffitte <paul.laffitte@enix.fr>
// Arthur Chaloin <arthur.chaloin@enix.fr>
// Alexandre Buisine <alexandre.buisine@enix.fr>
// Joe Skazinski <joseph.skazinski@seagate.com>

module github.com/vankosa/dothill-api-go/v2

go 1.12

require (
	github.com/enix/dothill-api-go/v2 v2.0.0
	github.com/gin-gonic/gin v1.7.2
	github.com/joho/godotenv v1.3.0
	github.com/onsi/gomega v1.13.0
	github.com/prometheus/client_golang v1.11.0
	k8s.io/klog v1.0.0
)
