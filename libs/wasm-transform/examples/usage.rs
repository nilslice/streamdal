use streamdal_wasm_transform::transform;
use streamdal_wasm_transform::transform::{TruncateOptions, TruncateType};

fn main() {
    println!("Overwrite example");
    overwrite();

    println!("Mask string example");
    mask_string();

    println!("Mask number example");
    mask_number();

    println!("Obfuscate string example");
    obfuscate_string();

    println!("Truncate string example");
    truncate_string();

    println!("Delete field example");
    delete_field();
}

fn overwrite() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = transform::Request {
        data: sample_json.into(),
        path: "hello".to_string(),
        value: r#""baz""#.to_string(),
        truncate_options: None,
    };

    let updated_json = transform::overwrite(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn mask_string() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = transform::Request {
        data: sample_json.into(),
        path: "hello".to_string(),
        value: "".to_string(),
        truncate_options: None,
    };

    let updated_json = transform::mask(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn mask_number() {
    let sample_json = r#"{"hello": 329328102938}"#;

    let req = transform::Request {
        path: "hello".to_string(),
        data: sample_json.into(),
        value: "".to_string(), // default
        truncate_options: None,
    };

    let updated_json = transform::mask(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn obfuscate_string() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = transform::Request {
        path: "hello".to_string(),
        data: sample_json.into(),
        value: "".to_string(), // default
        truncate_options: None,
    };

    let updated_json = transform::obfuscate(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn truncate_string() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = transform::Request {
        path: "hello".to_string(),
        data: sample_json.into(),
        value: "".to_string(), // default
        truncate_options: Some(TruncateOptions {
            length: 3,
            truncate_type: TruncateType::Chars,
        }),
    };

    let updated_json = transform::truncate(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn delete_field() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = transform::Request {
        path: "hello".to_string(),
        data: sample_json.into(),
        value: "".to_string(), // default
        truncate_options: None,
    };

    let updated_json = transform::delete(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}