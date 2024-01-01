extern crate libc;

use std::ffi::CString;

use libc::c_char;

#[no_mangle]
pub extern "C" fn hello_world(item: *mut c_char) -> *mut c_char {
    unsafe {
        let c_str =
            CString::new(format!("hello world {:?}", CString::from_raw(item))).expect("dont fail");
        c_str.into_raw()
    }
}

#[no_mangle]
pub extern "C" fn hello_world_free(item: *mut c_char) {
    unsafe {
        if item.is_null() {
            return;
        }
        let _ = CString::from_raw(item);
    };
}
