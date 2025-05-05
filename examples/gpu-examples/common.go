package gpuexamples

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

func LoadShader(device *sdl.GPUDevice, shaderFilename string, samplerCount, uniformBufferCount, storageBufferCount, storageTextureCount uint32) *sdl.GPUShader {
	// Auto-detect the shader stage from the file name for convenience
	var stage sdl.GPUShaderStage
	if sdl.Strstr(shaderFilename, ".vert") != "" {
		stage = sdl.GPUShaderStageVertex
	} else if sdl.Strstr(shaderFilename, ".frag") != "" {
		stage = sdl.GPUShaderStageFragment
	} else {
		sdl.Log("Invalid shader stage!")
		return nil
	}

	var fullPath, entryPoint string
	backendFormats := sdl.GetGPUShaderFormats(device)
	format := sdl.GPUShaderFormatInvalid

	wd, _ := os.Getwd()
	contentDir := filepath.Dir(wd)
	if backendFormats&sdl.GPUShaderFormatSPIRV != 0 {
		fullPath = fmt.Sprintf("%s/content/shaders/compiled/SPIRV/%s.spv", contentDir, shaderFilename)
		format = sdl.GPUShaderFormatSPIRV
		entryPoint = "main"
	} else if backendFormats&sdl.GPUShaderFormatMSL != 0 {
		fullPath = fmt.Sprintf("%s/content/shaders/compiled/MSL/%s.msl", contentDir, shaderFilename)
		format = sdl.GPUShaderFormatMSL
		entryPoint = "main0"
	} else if backendFormats&sdl.GPUShaderFormatDXIL != 0 {
		fullPath = fmt.Sprintf("%s/content/shaders/compiled/DXIL/%s.dxil", contentDir, shaderFilename)
		format = sdl.GPUShaderFormatDXIL
		entryPoint = "main"
	} else {
		sdl.Log("%s", "Unrecognized backend shader format!")
		return nil
	}

	var codeSize uint64
	code := sdl.LoadFile(fullPath, &codeSize)
	if code == nil {
		sdl.Log("Failed to load shader from disk! %s", fullPath)
		return nil
	}

	shaderInfo := sdl.GPUShaderCreateInfo{
		CodeSize:           codeSize,
		Code:               (*uint8)(code),
		EntryPoint:         convert.ToBytePtr(entryPoint),
		Format:             format,
		Stage:              stage,
		NumSamplers:        samplerCount,
		NumStorageTextures: uniformBufferCount,
		NumStorageBuffers:  storageBufferCount,
		NumUniformBuffers:  storageTextureCount,
	}

	shader := sdl.CreateGPUShader(device, &shaderInfo)
	if shader == nil {
		sdl.Log("Failed to create shader!")
		sdl.Free(code)
		return nil
	}

	sdl.Free(code)
	return shader
}
