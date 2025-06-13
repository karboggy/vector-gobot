// https://github.com/lvgl/lv_port_linux/blob/master/example/main.c

#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

#include "lvgl/lvgl.h"
// #include "lvgl/demos/lv_demos.h"
// #include "src/lib/driver_backends.h"

// #include <lvgl.h>
// #include <lvgl/demos/lv_demos.h>
// #include <lvgl/driver_backends.h>
// #include <lvgl/simulator_settings.h>

// extern simulator_settings_t settings;

int main(int argc, char **argv)
{

    /* Initialize LVGL. */
    lv_init();

    // framebuffer
    lv_display_t * disp = lv_linux_fbdev_create();
    lv_linux_fbdev_set_file(disp, "/dev/fb0");

    // background
    lv_obj_set_style_bg_color(lv_screen_active(), lv_color_hex(0x003a57), LV_PART_MAIN);

    // label
    lv_obj_t * label = lv_label_create(lv_screen_active());
    lv_label_set_text(label, "Hello world");
    lv_obj_set_style_text_color(lv_screen_active(), lv_color_hex(0xffffff), LV_PART_MAIN);
    lv_obj_align(label, LV_ALIGN_CENTER, 0, 0);

    while(1) {
        uint32_t ms = lv_timer_handler();
        usleep(ms*1000);
    }

    // settings.window_width = 800;
    // settings.window_height = 480;

    /* Initialize the configured default backend */
    // driver_backends_register();
    // if (driver_backends_init_backend("WAYLAND") == -1) {
    //     fprintf(stderr, "Failed to initialize wayland backend\n");
    //     exit(EXIT_FAILURE);
    // }

    // /*Create a Demo*/
    // lv_demo_widgets();
    // lv_demo_widgets_start_slideshow();

    // /* Enter the run loop - does not return */
    // driver_backends_run_loop();

    return 0;
}