<?php
// Create a stream
$opts = array(
    'http'=>array(
      'method'=>"GET",
      'content' => "sleep 10"
    )
  );
  
  $context = stream_context_create($opts);
  
  // Open the file using the HTTP headers set above
  $file = file_get_contents('http://127.0.0.1:8585/exe', false, $context);
  echo "....".$file."\r\n";
?>