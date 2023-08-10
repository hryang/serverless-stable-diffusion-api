package module

// GetImagesRequest defines model for GetImagesRequest.
type GetImagesRequest struct {
	TaskId []string `json:"taskId"`
}

// Image defines model for Image.
type Image struct {
	Images     []string                `json:"images"`
	Info       *map[string]interface{} `json:"info,omitempty"`
	Parameters *map[string]interface{} `json:"parameters,omitempty"`
	TaskId     string                  `json:"taskId"`
}

// ImageResponse defines model for ImageResponse.
type ImageResponse struct {
	Data   []Image `json:"data"`
	ErrMsg *string `json:"errMsg,omitempty"`
}

// Img2ImgRequest defines model for Img2ImgRequest.
type Img2ImgRequest struct {
	AlwaysonScripts                   *map[string]interface{} `json:"alwayson_scripts,omitempty"`
	BatchSize                         *int32                  `json:"batch_size,omitempty"`
	CfgScale                          *int32                  `json:"cfg_scale,omitempty"`
	DenoisingStrength                 *float32                `json:"denoising_strength,omitempty"`
	DoNotSaveGrid                     *bool                   `json:"do_not_save_grid,omitempty"`
	DoNotSaveSamples                  *bool                   `json:"do_not_save_samples,omitempty"`
	Eta                               *int32                  `json:"eta,omitempty"`
	Height                            *int32                  `json:"height,omitempty"`
	ImageCfgScale                     *int32                  `json:"image_cfg_scale,omitempty"`
	IncludeInitImages                 *bool                   `json:"include_init_images,omitempty"`
	InitImages                        *[]string               `json:"init_images,omitempty"`
	InitialNoiseMultiplier            *int32                  `json:"initial_noise_multiplier,omitempty"`
	InpaintFullRes                    *bool                   `json:"inpaint_full_res,omitempty"`
	InpaintFullResPadding             *int32                  `json:"inpaint_full_res_padding,omitempty"`
	InpaintingFill                    *int32                  `json:"inpainting_fill,omitempty"`
	InpaintingMaskInvert              *int32                  `json:"inpainting_mask_invert,omitempty"`
	Mask                              *string                 `json:"mask,omitempty"`
	MaskBlur                          *int32                  `json:"mask_blur,omitempty"`
	MaskBlurX                         *int32                  `json:"mask_blur_x,omitempty"`
	MaskBlurY                         *int32                  `json:"mask_blur_y,omitempty"`
	NIter                             *int32                  `json:"n_iter,omitempty"`
	NegativePrompt                    *string                 `json:"negative_prompt,omitempty"`
	OverrideSettings                  *map[string]interface{} `json:"override_settings,omitempty"`
	OverrideSettingsRestoreAfterwards *bool                   `json:"override_settings_restore_afterwards,omitempty"`
	Prompt                            *string                 `json:"prompt,omitempty"`
	ResizeMode                        *int32                  `json:"resize_mode,omitempty"`
	RestoreFaces                      *bool                   `json:"restore_faces,omitempty"`
	SChurn                            *int32                  `json:"s_churn,omitempty"`
	SMinUncond                        *int32                  `json:"s_min_uncond,omitempty"`
	SNoise                            *int32                  `json:"s_noise,omitempty"`
	STmax                             *int32                  `json:"s_tmax,omitempty"`
	STmin                             *int32                  `json:"s_tmin,omitempty"`
	SamplerIndex                      *string                 `json:"sampler_index,omitempty"`
	SamplerName                       *string                 `json:"sampler_name,omitempty"`
	SaveDir                           *string                 `json:"save_dir,omitempty"`
	SaveImages                        *bool                   `json:"save_images,omitempty"`
	ScriptArgs                        *[]string               `json:"script_args,omitempty"`
	ScriptName                        *string                 `json:"script_name,omitempty"`
	SdVae                             string                  `json:"sd_vae"`
	Seed                              *int32                  `json:"seed,omitempty"`
	SeedResizeFromH                   *int32                  `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW                   *int32                  `json:"seed_resize_from_w,omitempty"`
	SendImages                        *bool                   `json:"send_images,omitempty"`
	StableDiffusionModel              string                  `json:"stable_diffusion_model"`
	Steps                             *int32                  `json:"steps,omitempty"`
	Styles                            *[]string               `json:"styles,omitempty"`
	Subseed                           *int32                  `json:"subseed,omitempty"`
	SubseedStrength                   *int32                  `json:"subseed_strength,omitempty"`
	Tiling                            *bool                   `json:"tiling,omitempty"`
	Width                             *int32                  `json:"width,omitempty"`
}

// Img2ImgResponse defines model for Img2ImgResponse.
type Img2ImgResponse struct {
	ErrMsg *string `json:"err_msg,omitempty"`
	Status string  `json:"status"`
	TaskId string  `json:"taskId"`
}

// ProgressRespons defines model for ProgressRespons.
type ProgressRespons struct {
	CurrentImage string                  `json:"currentImage"`
	EatRelative  *float32                `json:"eatRelative,omitempty"`
	Progress     float32                 `json:"progress"`
	State        *map[string]interface{} `json:"state,omitempty"`
	Status       *string                 `json:"status,omitempty"`
	TaskId       string                  `json:"taskId"`
}

// Txt2ImgRequest defines model for Txt2ImgRequest.
type Txt2ImgRequest struct {
	AlwaysonScripts                   *map[string]interface{}   `json:"alwayson_scripts,omitempty"`
	BatchSize                         *int32                    `json:"batch_size,omitempty"`
	CfgScale                          *int32                    `json:"cfg_scale,omitempty"`
	DenoisingStrength                 *float32                  `json:"denoising_strength,omitempty"`
	DoNotSaveGrid                     *bool                     `json:"do_not_save_grid,omitempty"`
	DoNotSaveSamples                  *bool                     `json:"do_not_save_samples,omitempty"`
	EnableHr                          *bool                     `json:"enable_hr,omitempty"`
	Eta                               *int32                    `json:"eta,omitempty"`
	FirstphaseHeight                  *int32                    `json:"firstphase_height,omitempty"`
	FirstphaseWidth                   *int32                    `json:"firstphase_width,omitempty"`
	Height                            *int32                    `json:"height,omitempty"`
	HrNegativePrompt                  *string                   `json:"hr_negative_prompt,omitempty"`
	HrPrompt                          *string                   `json:"hr_prompt,omitempty"`
	HrResizeX                         *int32                    `json:"hr_resize_x,omitempty"`
	HrResizeY                         *int32                    `json:"hr_resize_y,omitempty"`
	HrSamplerName                     *string                   `json:"hr_sampler_name,omitempty"`
	HrScale                           *int32                    `json:"hr_scale,omitempty"`
	HrSecondPassSteps                 *int32                    `json:"hr_second_pass_steps,omitempty"`
	HrUpscaler                        *string                   `json:"hr_upscaler,omitempty"`
	NIter                             *int32                    `json:"n_iter,omitempty"`
	NegativePrompt                    *string                   `json:"negative_prompt,omitempty"`
	OverrideSettings                  *map[string]interface{}   `json:"override_settings,omitempty"`
	OverrideSettingsRestoreAfterwards *bool                     `json:"override_settings_restore_afterwards,omitempty"`
	Prompt                            string                    `json:"prompt"`
	RestoreFaces                      *bool                     `json:"restore_faces,omitempty"`
	SChurn                            *int32                    `json:"s_churn,omitempty"`
	SMinUncond                        *int32                    `json:"s_min_uncond,omitempty"`
	SNoise                            *int32                    `json:"s_noise,omitempty"`
	STmax                             *int32                    `json:"s_tmax,omitempty"`
	STmin                             *int32                    `json:"s_tmin,omitempty"`
	SamplerIndex                      *string                   `json:"sampler_index,omitempty"`
	SamplerName                       *string                   `json:"sampler_name,omitempty"`
	SaveDir                           *string                   `json:"save_dir,omitempty"`
	SaveImages                        *bool                     `json:"save_images,omitempty"`
	ScriptArgs                        *[]map[string]interface{} `json:"script_args,omitempty"`
	ScriptName                        *string                   `json:"script_name,omitempty"`
	SdVae                             string                    `json:"sd_vae"`
	Seed                              *int32                    `json:"seed,omitempty"`
	SeedResizeFromH                   *int32                    `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW                   *int32                    `json:"seed_resize_from_w,omitempty"`
	SendImages                        *bool                     `json:"send_images,omitempty"`
	StableDiffusionModel              string                    `json:"stable_diffusion_model"`
	Steps                             *int32                    `json:"steps,omitempty"`
	Styles                            *[]string                 `json:"styles,omitempty"`
	Subseed                           *int32                    `json:"subseed,omitempty"`
	SubseedStrength                   *int32                    `json:"subseed_strength,omitempty"`
	Tiling                            *bool                     `json:"tiling,omitempty"`
	Width                             *int32                    `json:"width,omitempty"`
}

// Txt2ImgResponse defines model for Txt2ImgResponse.
type Txt2ImgResponse struct {
	ErrMsg *string `json:"err_msg,omitempty"`
	Status string  `json:"status"`
	TaskId string  `json:"taskId"`
}

// GetCancelParams defines parameters for GetCancel.
type GetCancelParams struct {
	// TaskId task id
	TaskId string `form:"taskId" json:"taskId"`
}

// GetGetResultParams defines parameters for GetGetResult.
type GetGetResultParams struct {
	// TaskId task id
	TaskId string `form:"taskId" json:"taskId"`
}

// GetProgressParams defines parameters for GetProgress.
type GetProgressParams struct {
	// TaskId task id
	TaskId string `form:"taskId" json:"taskId"`
}

// PostGetResultJSONRequestBody defines body for PostGetResult for application/json ContentType.
type PostGetResultJSONRequestBody = GetImagesRequest

// Img2ImgJSONRequestBody defines body for Img2Img for application/json ContentType.
type Img2ImgJSONRequestBody = Img2ImgRequest

// Txt2ImgJSONRequestBody defines body for Txt2Img for application/json ContentType.
type Txt2ImgJSONRequestBody = Txt2ImgRequest
