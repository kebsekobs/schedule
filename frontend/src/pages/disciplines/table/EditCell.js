import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useState } from "react";
import EditModal from "../modals/editModal.js";
import styles from '../../shared/style/table.module.css';
import { DeleteModal } from "../modals/deleteModal.js";

export function EditCell(props) {
  const id = props.props.original.id;
  const original = props.props.original;

  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDeleteModalOpen, setisDeleteModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteDiscipline() {
    setisDeleteModalOpen(!isDeleteModalOpen);
  }

  return (
    <div style={{padding: '5px'}}>
      <div
        className={styles["btns-wrapper"]}
        onClick={() => toggleEditModal(id)}
      >
        <p>Изменить</p>
        <Pencil1Icon />
      </div>
      <div className={styles["btns-wrapper"]} onClick={deleteDiscipline}>
        <p style={{ color: "var(--warning-color)" }}>Удалить</p>
        <TrashIcon />
      </div>
      {isEditModalOpen && (
        <EditModal
          toggleModal={toggleEditModal}
          isOpen={isEditModalOpen}
          original={original}
        />
      )}
            {isDeleteModalOpen && (
        <DeleteModal
        deleteDiscipline={deleteDiscipline}
          isOpen={isDeleteModalOpen}
          original={original}
        />
      )}
    </div>
  );
}
