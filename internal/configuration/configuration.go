package configuration

type Region struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    struct {
		Lat  float64 `json:"lat"`
		Name string  `json:"name"`
		Long float64 `json:"long"`
	} `json:"location"`
	Services struct {
		Type                   string `json:"type"`
		SingleNamespace        bool   `json:"singleNamespace"`
		AllowNamespaceCreation bool   `json:"allowNamespaceCreation"`
		NamespaceLabels        struct {
		} `json:"namespaceLabels"`
		NamespaceAnnotations struct {
		} `json:"namespaceAnnotations"`
		UserNamespace        bool   `json:"userNamespace"`
		NamespacePrefix      string `json:"namespacePrefix"`
		GroupNamespacePrefix string `json:"groupNamespacePrefix"`
		UsernamePrefix       string `json:"usernamePrefix"`
		GroupPrefix          string `json:"groupPrefix"`
		AuthenticationMode   string `json:"authenticationMode"`
		Expose               struct {
			Domain                string `json:"domain"`
			IngressClassName      string `json:"ingressClassName"`
			UseDefaultCertificate bool   `json:"useDefaultCertificate"`
			Annotations           struct {
			} `json:"annotations"`
			Ingress bool `json:"ingress"`
			Route   bool `json:"route"`
		} `json:"expose"`
		Monitoring struct {
			URLPattern string `json:"URLPattern"`
		} `json:"monitoring"`
		AllowedURIPattern string `json:"allowedURIPattern"`
		Quotas            struct {
			Enabled               bool `json:"enabled"`
			UserEnabled           bool `json:"userEnabled"`
			GroupEnabled          bool `json:"groupEnabled"`
			AllowUserModification bool `json:"allowUserModification"`
			Default               struct {
				RequestsStorage      string `json:"requests.storage"`
				CountPods            int    `json:"count/pods"`
				RequestsNvidiaComGpu int    `json:"requests.nvidia.com/gpu"`
			} `json:"default"`
		} `json:"quotas"`
		DefaultConfiguration struct {
			NetworkPolicy bool `json:"networkPolicy"`
			From          []struct {
				IPBlock struct {
					Cidr string `json:"cidr"`
				} `json:"ipBlock"`
			} `json:"from"`
			Tolerations  []any `json:"tolerations"`
			StartupProbe struct {
				FailureThreshold    int `json:"failureThreshold"`
				InitialDelaySeconds int `json:"initialDelaySeconds"`
				PeriodSeconds       int `json:"periodSeconds"`
				SuccessThreshold    int `json:"successThreshold"`
				TimeoutSeconds      int `json:"timeoutSeconds"`
			} `json:"startupProbe"`
			Kafka struct {
				TopicName string `json:"topicName"`
				URL       string `json:"URL"`
			} `json:"kafka"`
			Sliders struct {
				CPU struct {
					SliderMin  int    `json:"sliderMin"`
					SliderMax  int    `json:"sliderMax"`
					SliderStep int    `json:"sliderStep"`
					SliderUnit string `json:"sliderUnit"`
				} `json:"cpu"`
				Memory struct {
					SliderMin  int    `json:"sliderMin"`
					SliderMax  int    `json:"sliderMax"`
					SliderStep int    `json:"sliderStep"`
					SliderUnit string `json:"sliderUnit"`
				} `json:"memory"`
				Gpu struct {
					SliderMin  int    `json:"sliderMin"`
					SliderMax  int    `json:"sliderMax"`
					SliderStep int    `json:"sliderStep"`
					SliderUnit string `json:"sliderUnit"`
				} `json:"gpu"`
				Disk struct {
					SliderMin  int    `json:"sliderMin"`
					SliderMax  int    `json:"sliderMax"`
					SliderStep int    `json:"sliderStep"`
					SliderUnit string `json:"sliderUnit"`
				} `json:"disk"`
			} `json:"sliders"`
			Resources struct {
				CPURequest    string `json:"cpuRequest"`
				CPULimit      string `json:"cpuLimit"`
				MemoryRequest string `json:"memoryRequest"`
				MemoryLimit   string `json:"memoryLimit"`
				Disk          string `json:"disk"`
				Gpu           string `json:"gpu"`
			} `json:"resources"`
			Ipprotection bool `json:"ipprotection"`
		} `json:"defaultConfiguration"`
		K8SPublicEndpoint struct {
			OidcConfiguration struct {
				IssuerURI string `json:"issuerURI"`
				ClientID  string `json:"clientID"`
			} `json:"oidcConfiguration"`
			URL string `json:"URL"`
		} `json:"k8sPublicEndpoint"`
		CustomInitScript struct {
		} `json:"customInitScript"`
		CustomValues struct {
		} `json:"customValues"`
	} `json:"services"`
	Data struct {
		Atlas struct {
			OidcConfiguration struct {
				IssuerURI string `json:"issuerURI"`
				ClientID  string `json:"clientID"`
			} `json:"oidcConfiguration"`
			URL string `json:"URL"`
		} `json:"atlas"`
		S3 struct {
			Region          string `json:"region"`
			PathStyleAccess bool   `json:"pathStyleAccess"`
			Sts             struct {
				DurationSeconds   int `json:"durationSeconds"`
				OidcConfiguration struct {
					ClientID string `json:"clientID"`
				} `json:"oidcConfiguration"`
			} `json:"sts"`
			WorkingDirectory struct {
				BucketMode            string `json:"bucketMode"`
				BucketNamePrefix      string `json:"bucketNamePrefix"`
				BucketNamePrefixGroup string `json:"bucketNamePrefixGroup"`
			} `json:"workingDirectory"`
			URL string `json:"URL"`
		} `json:"S3"`
	} `json:"data"`
	Vault struct {
		KvEngine string `json:"kvEngine"`
		Role     string `json:"role"`
		AuthPath string `json:"authPath"`
		URL      string `json:"URL"`
	} `json:"vault"`
	Git struct {
		OidcConfiguration struct {
			IssuerURI string `json:"issuerURI"`
			ClientID  string `json:"clientID"`
		} `json:"oidcConfiguration"`
		URL string `json:"URL"`
	} `json:"git"`
}

type Configuration struct {
	Authentication Authentication
	RootPath       string
	Regions        []Region
}

type Authentication struct {
	IssuerURI string
	Audience  string
}
