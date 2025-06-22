#!/bin/bash

mkdir -p samples 
# Create the 'samples' directory if it doesn't exist


# Array of URLs to download
image_urls=(
    "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/boots-image-horizontal.png"
    "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/boots-image-vertical.png"
    "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/boots-video-horizontal.mp4"
    "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/boots-video-vertical.mp4"
    "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/is-bootdev-for-you.pdf"
)

# Loop through each URL in the array
for url in "${image_urls[@]}"; do # Extract the file name from the URL 
  file_name=$(basename "$url") # Download the file and save it in the 'samples' directory
  curl -sSfL -o "samples/$file_name" "$url"
done
