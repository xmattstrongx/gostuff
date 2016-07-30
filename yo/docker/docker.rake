PROJECT_ROOT = Dir.pwd
DOCKER_DIR = "#{PROJECT_ROOT}/docker"
SERVICES_DIR = "#{DOCKER_DIR}/services"
DOCKER_SERVICES = `ls #{SERVICES_DIR}`.split()
OUTPUT_DIR = "#{DOCKER_DIR}/bin"
BUILD_FILE = 'Rakefile'
IMAGE_PREFIX = 'generic'
REGISTRY_URL = 'ps-reg.mia.ucloud.int'
OUTPUTS = DOCKER_SERVICES.map{|name| File.join(OUTPUT_DIR, name)}
CLEAN.include(OUTPUT_DIR)

namespace :docker do
# Create a build task for each service
DOCKER_SERVICES.each do |service|
  desc "Build #{service}"
  task service => ["docker:check_docker", "docker:check_conn", :godep] do
    # build any artifacts for this service
    # run the build file for this service if it exists
    service_dir = "#{SERVICES_DIR}/#{service}"
    build_filename = "#{service_dir}/#{BUILD_FILE}"
    output_dir = "#{OUTPUT_DIR}/#{service}"
    sh "mkdir -p #{output_dir}"
    if File.exist?(build_filename)
      sh "rake -f #{service_dir}/#{BUILD_FILE} default[#{PROJECT_ROOT},#{output_dir}]"
    end
    # now build the image for this service
    build_service_image(service, service_dir, output_dir)
  end

end
desc "Build all services"
task :build => DOCKER_SERVICES.each {|service| service}

desc "Build all services concurrently"
multitask :multibuild => DOCKER_SERVICES.each {|service| service}

task :godep do
    sh "go get -v github.com/tools/godep"
end

def build_service_image(service_name, service_dir, output_dir)

  # copy all of the files except the builder file
  files = `ls #{service_dir}`
    .split().select{|file| file != BUILD_FILE}
    .each{|file| sh "cp #{service_dir}/#{file} #{output_dir}/#{file}"}

  # build the docker image
  sh "docker build -t #{IMAGE_PREFIX}-#{service_name} #{output_dir}"
end

task :check_docker do
  fail_if_cmd_fails('which docker', "You don't seem to have docker installed. Stopping...")
end

task :check_conn do
  fail_if_cmd_fails('docker ps', "Something appears to be wrong with your Docker connection. Stopping...")
end

task :check_compose do
  fail_if_cmd_fails('which docker-compose', "You don't seem to have docker-compose installed. Stopping...")
end

def fail_if_cmd_fails(cmd, msg)
  output = `#{cmd}`
  if $?.exitstatus != 0
    puts output
    fail msg
  end
  return output
end

desc "Start the generic project"
task :start => ["docker:check_docker", "docker:check_conn", "docker:check_compose"] do
  sh "docker-compose up -d"
end

task :push => ["docker:check_docker", "docker:check_conn"] do
  DOCKER_SERVICES.each do |service|
    image_name = "#{IMAGE_PREFIX}-#{service}"
    sh "docker tag -f #{image_name} #{REGISTRY_URL}/#{image_name}"
    sh "docker push #{REGISTRY_URL}/#{image_name}"
  end
end
end
