# Located in /etc/puppetlabs/puppet/modules/golang_reports/lib/puppet/reports/golang_reports.rb
require 'puppet'
require 'net/http'
require 'uri'
require 'json'

Puppet::Reports.register_report(:golang_reports) do
  desc 'Process reports via the golang_reports API.'

  def process
    p_report = {
      'certname' => host,
      'environment' => environment,
      'status' => status,
      'time' => time,
      'transaction_uuid' => transaction_uuid,
    }

    begin
      send_report(p_report)
      Puppet.info 'Sending Report to Golang Reports API.'
    rescue StandardError => e
      Puppet.err "Could not send report to Golang Reports API: #{e}\n#{e.backtrace}"
    end
  end

  def send_report(body)
    reports_url = 'http://example.hostname.com:2754/reports/add'
    @uri = URI.parse(reports_url)
    http = Net::HTTP.new(@uri.host, @uri.port)
    http.verify_mode = OpenSSL::SSL::VERIFY_NONE
    req = Net::HTTP::Post.new(@uri.path.to_str)
    req.add_field('Content-Type', 'application/json')
    req.content_type = 'application/json'
    req.body = body.to_json
    http.request(req)
  end
end
