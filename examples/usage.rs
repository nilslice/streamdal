use streamdal_transform::transform::Request;

fn main() {
    println!("Overwrite example");
    overwrite();

    println!("Mask string example");
    mask_string();

    println!("Mask number example");
    mask_number();

    println!("Obfuscate string example");
    obfuscate_string()
}

fn overwrite() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = Request {
        data: sample_json.into(),
        path: "hello".to_string(),
        value: r#""baz""#.to_string(),
    };

    let updated_json = streamdal_transform::transform::overwrite(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn mask_string() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = Request {
        data: sample_json.into(),
        path: "hello".to_string(),
        value: "".to_string(),
    };

    let updated_json = streamdal_transform::transform::mask(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn mask_number() {
    let sample_json = r#"{"hello": 329328102938}"#;

    let req = Request {
        path: "hello".to_string(),
        data: sample_json.into(),
        value: "".to_string(), // default
    };

    let updated_json = streamdal_transform::transform::mask(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}

fn obfuscate_string() {
    let sample_json = r#"{"hello": "world"}"#;

    let req = Request {
        path: "hello".to_string(),
        data: sample_json.into(),
        value: "".to_string(), // default
    };

    let updated_json = streamdal_transform::transform::obfuscate(&req).unwrap();

    println!(
        "Input JSON: {} || Result JSON: {}",
        sample_json, updated_json,
    )
}
