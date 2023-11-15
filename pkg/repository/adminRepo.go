package repository

import (
	"fmt"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
	"github.com/Nishad4140/ecommerce_project/pkg/repository/controllers"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

//-------------------------- Login --------------------------//

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}

//-------------------------- Report-User --------------------------//

func (c *adminDatabase) ReportUser(reason helper.ReportData, adminId int) error {
	tx := c.DB.Begin()

	var Exists bool
	if err := tx.Raw("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", reason.UserId).Scan(&Exists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !Exists {
		tx.Rollback()
		return fmt.Errorf("no such a user")
	}
	if err := tx.Exec("UPDATE users SET report_count=report_count+1 WHERE id = ?", reason.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("INSERT INTO user_report_infos (users_id, reason_for_reporting, reported_at, reported_by) VALUES (?, ?, NOW(), ?)", reason.UserId, reason.Reason, adminId).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

//-------------------------- Show-Single-User --------------------------//

func (c *adminDatabase) ShowUser(userID int) (response.UserDetails, error) {
	var userData response.UserDetails
	qury := `SELECT users.name,
			 users.email, 
			 users.mobile,
			 users.report_count,  
			 users.is_blocked, 
			 block_infos.blocked_by,
			 block_infos.blocked_at,
			 block_infos.reason_for_blocking 
			 FROM users as users 
			 FULL OUTER JOIN user_report_infos as report_infos ON users.id = report_infos.users_id
			 FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id
			 WHERE users.id = $1;`

	err := c.DB.Raw(qury, userID).Scan(&userData).Error
	if err != nil {
		return response.UserDetails{}, err
	}
	if userData.Email == "" {
		return response.UserDetails{}, fmt.Errorf("no such user")
	}
	return userData, nil
}

//-------------------------- Show-All-Users --------------------------//

func (c *adminDatabase) ShowAllUser() ([]response.UserDetails, error) {
	var userDatas []response.UserDetails

	getUsers := `SELECT users.name,
				users.email, 
				users.mobile,
				users.report_count,  
				users.is_blocked, 
				block_infos.blocked_by,
				block_infos.blocked_at,
				block_infos.reason_for_blocking 
				FROM users as users 
				FULL OUTER JOIN user_report_infos as report_infos ON users.id = report_infos.users_id
				FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id;`

	err := c.DB.Raw(getUsers).Scan(&userDatas).Error
	return userDatas, err
}

//-------------------------- Dashboard --------------------------//

func (c *adminDatabase) GetDashBoard(reports helper.ReportParams) (response.DashBoard, error) {
	tx := c.DB.Begin()
	var dashBoard response.DashBoard
	getDasheBoard := `SELECT SUM(oi.quantity*oi.price)as Total_Revenue,
			SUM (oi.quantity)as Total_Products_Selled,
			COUNT(DISTINCT o.id)as Total_Orders FROM orders o
			JOIN order_items oi on o.id=oi.orders_id`

	getTotalUsers := `SELECT COUNT(id)AS TotalUsers FROM users`

	if reports.Status != "" {
		var status domain.OrderStatus
		err := c.DB.Raw("SELECT * FROM order_statuses WHERE status=?", reports.Status).Scan(&status).Error
		if err != nil {
			return dashBoard, fmt.Errorf("error scaning the order status table")
		}
		getDasheBoard = fmt.Sprintf("%s WHERE o.order_status_id=%d", getDasheBoard, status.Id)
		// getTotalUsers = fmt.Sprintf("%s WHERE o.order_status_id=%d", getTotalUsers, status.Id)
	} else {
		getDasheBoard = fmt.Sprintf("%s WHERE o.order_status_id is not null", getDasheBoard)
		// getTotalUsers = fmt.Sprintf("%s WHERE o.order_status_id is not null", getTotalUsers)
	}

	if reports.Day != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getDasheBoard = fmt.Sprintf("%s AND o.order_date::date='%s'", getDasheBoard, date)
		getTotalUsers = fmt.Sprintf("%s WHERE created_at::date='%s'", getDasheBoard, date)
	} else if reports.Week != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getDasheBoard = fmt.Sprintf("%s AND o.order_date BETWEEN %s", getDasheBoard, date)
		getTotalUsers = fmt.Sprintf("%s WHERE created_at BETWEEN %s", getDasheBoard, date)
	} else if reports.Month != 0 && reports.Year != 0 {
		getDasheBoard = fmt.Sprintf("%s AND EXTRACT(YEAR FROM order_date) = %d AND EXTRACT(MONTH FROM order_date) = %d", getDasheBoard, reports.Year, reports.Month)
		getTotalUsers = fmt.Sprintf("%s WHERE EXTRACT(YEAR FROM created_at) = %d AND EXTRACT(MONTH FROM created_at) = %d", getTotalUsers, reports.Year, reports.Month)
	} else if reports.Date1 != "" && reports.Date2 != "" {
		getDasheBoard = fmt.Sprintf("%s AND o.order_date BETWEEN '%s 00:00:00'::timestamp AND '%s 23:59:59'::timestamp", getDasheBoard, reports.Date1, reports.Date2)
		getTotalUsers = fmt.Sprintf("%s WHERE created_at BETWEEN '%s 00:00:00'::timestamp AND '%s 23:59:59'::timestamp", getTotalUsers, reports.Date1, reports.Date2)
	} else if reports.Year != 0 {
		getDasheBoard = fmt.Sprintf("%s AND EXTRACT ( YEAR FROM order_date) = %d", getDasheBoard, reports.Year)
		getTotalUsers = fmt.Sprintf("%s WHERE EXTRACT ( YEAR FROM created_at) = %d", getTotalUsers, reports.Year)
	}

	if err := tx.Raw(getDasheBoard).Scan(&dashBoard).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}
	// getDasheBoard := `SELECT SUM(quantity*price)as Total_Revenue,
	// 		SUM (quantity)as Total_Products_Selled FROM order_items`
	// if err := tx.Raw(getDasheBoard).Scan(&dashBoard).Error; err != nil {
	// 	tx.Rollback()
	// 	return response.DashBoard{}, err
	// }

	// getOrderNo := `SELECT COUNT(id)FROM orders WHERE order_status_id=$1`
	// if err := tx.Raw(getOrderNo, 1).Scan(&dashBoard.TotalOrders).Error; err != nil {
	// 	tx.Rollback()
	// 	return response.DashBoard{}, err
	// }

	if err := tx.Raw(getTotalUsers).Scan(&dashBoard.TotalUsers).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}
	return dashBoard, nil
}

func (c *adminDatabase) ViewSalesReport(reports helper.ReportParams) ([]response.SalesReport, error) {

	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id`

	if reports.Status != "" {
		var status domain.OrderStatus
		err := c.DB.Raw("SELECT * FROM order_statuses WHERE status=?", reports.Status).Scan(&status).Error
		if err != nil {
			return sales, fmt.Errorf("error scaning the order status table")
		}
		getReports = fmt.Sprintf("%s WHERE o.order_status_id=%d", getReports, status.Id)
	} else {
		getReports = fmt.Sprintf("%s WHERE o.order_status_id is not null", getReports)
	}

	if reports.Day != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getReports = fmt.Sprintf("%s AND o.order_date::date='%s'", getReports, date)
	} else if reports.Week != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getReports = fmt.Sprintf("%s AND o.order_date BETWEEN %s", getReports, date)
	} else if reports.Month != 0 && reports.Year != 0 {
		getReports = fmt.Sprintf("%s AND EXTRACT(YEAR FROM order_date) = %d AND EXTRACT(MONTH FROM order_date) = %d", getReports, reports.Year, reports.Month)
	} else if reports.Date1 != "" && reports.Date2 != "" {
		getReports = fmt.Sprintf("%s AND o.order_date BETWEEN '%s 00:00:00'::timestamp AND '%s 23:59:59'::timestamp", getReports, reports.Date1, reports.Date2)
	} else if reports.Year != 0 {
		getReports = fmt.Sprintf("%s AND EXTRACT ( YEAR FROM order_date) = %d", getReports, reports.Year)
	}
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err
}
