package sdl

// [VulkanLoadLibrary] dynamically loads the Vulkan loader library.
//
// [VulkanLoadLibrary]: https://wiki.libsdl.org/SDL3/SDL_Vulkan_LoadLibrary
// func VulkanLoadLibrary(path *byte) bool {
// 	return sdlVulkanLoadLibrary(path)
// }

// [VulkanGetVkGetInstanceProcAddr] gets the address of the `vkGetInstanceProcAddr` function.
//
// [VulkanGetVkGetInstanceProcAddr]: https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetVkGetInstanceProcAddr
// func VulkanGetVkGetInstanceProcAddr() FunctionPointer {
// 	return sdlVulkanGetVkGetInstanceProcAddr()
// }

// [VulkanUnloadLibrary] unloads the Vulkan library previously loaded by [Vulkan_LoadLibrary].
//
// [VulkanUnloadLibrary]: https://wiki.libsdl.org/SDL3/SDL_Vulkan_UnloadLibrary
// func VulkanUnloadLibrary() {
// 	sdlVulkanUnloadLibrary()
// }

// [VulkanGetInstanceExtensions] gets the Vulkan instance extensions needed for vkCreateInstance.
//
// [VulkanGetInstanceExtensions]: https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetInstanceExtensions
// func VulkanGetInstanceExtensions(count *uint32) *byte {
// 	return sdlVulkanGetInstanceExtensions(count)
// }

// [VulkanCreateSurface] creates a Vulkan rendering surface for a window.
//
// [VulkanCreateSurface]: https://wiki.libsdl.org/SDL3/SDL_Vulkan_CreateSurface
// func VulkanCreateSurface(window *Window, instance uintptr, allocator uintptr, surface uintptr) bool {
// 	return sdlVulkanCreateSurface(window, instance, allocator, surface)
// }

// [VulkanDestroySurface] destroys the Vulkan rendering surface of a window.
//
// [VulkanDestroySurface]: https://wiki.libsdl.org/SDL3/SDL_Vulkan_DestroySurface
// func VulkanDestroySurface(instance uintptr, surface uintptr, allocator uintptr) {
// 	sdlVulkanDestroySurface(instance, surface, allocator)
// }

// [VulkanGetPresentationSupport] queries support for presentation via a given physical device and queue.
//
// [VulkanGetPresentationSupport]: https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetPresentationSupport
// func VulkanGetPresentationSupport(instance uintptr, physicalDevice uintptr, queueFamilyIndex uint32) bool {
// 	return sdlVulkanGetPresentationSupport(instance, physicalDevice, queueFamilyIndex)
// }
